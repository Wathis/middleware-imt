import React from 'react';
import clsx from 'clsx';
import { withStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Moyennes from './components/Moyennes';
import SelectAirport from './components/SelectAirport';
import axios from 'axios';
import Chart from './components/Chart';

const styles = theme => ({
  root: {
    display: 'flex',
  },
  menuButton: {
    marginRight: 36,
  },
  title: {
    flexGrow: 1,
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column',
  },
  fixedHeight: {
    height: 240,
  },
  centerText: {
    textAlign: "center",
  },
});

class App extends React.Component {

  constructor() {
    super();
    this.state = {
      airportId: "NTE",
      airportList: ["NTE", "CDG", "LYS"],
      measures: [],
    }
    this.handleChangeSelectedAirport = this.handleChangeSelectedAirport.bind(this)
  }

  handleChangeSelectedAirport(airport) {
    this.setState({ airportId: airport })
  }

  componentDidMount() {
    this.loadGraphData()
  }

  render() {
    const { classes } = this.props;
    const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);

    return (
      <div className={classes.root}>
        <CssBaseline />
        <AppBar position="absolute" className={classes.appBar}>
          <Toolbar className={classes.toolbar}>
            <Typography component="h1" variant="h6" color="inherit" noWrap className={classes.title}>
              Visualisation des mesures
            </Typography>
          </Toolbar>
        </AppBar>
        <main className={classes.content}>
          <div className={classes.appBarSpacer} />
          <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={3}>
              {/* Moyenne des temperatures de la journée */}
              <Grid item xs={12} md={12} lg={12}>
                <Moyennes />
              </Grid>
              <Grid item xs={12}>
                <SelectAirport 
                  airportList={["NTE", "CDG", "LYS"]}
                  onAirportChange={this.handleChangeSelectedAirport}
                />
              </Grid>

              {/* Chart */}
              <Grid item xs={12} md={12} lg={12}>
                <Paper className={fixedHeightPaper}>
                  <Chart
                    title="Evolution des relevés de température pour l'aéroport de "
                    data={this.state.measures}
                    measure="TEMP"
                    airportId={this.state.airportId}
                    xAxis="Température (°c)"
                  />
                </Paper>
              </Grid>
              <Grid item xs={12} md={12} lg={12}>
                <Paper className={fixedHeightPaper}>
                  <Chart
                    title="Evolution des relevés de pression pour l'aéroport de "
                    data={this.state.measures}
                    measure="PRESSURE"
                    airportId={this.state.airportId}
                    xAxis="Pression (Pa)"
                  />
                </Paper>
              </Grid>
              <Grid item xs={12} md={12} lg={12}>
                <Paper className={fixedHeightPaper}>
                  <Chart
                    title="Evolution des relevés de vent pour l'aéroport de "
                    data={this.state.measures}
                    measure="WIND"
                    airportId={this.state.airportId}
                    xAxis="Vent (km.h)"
                  />
                </Paper>
              </Grid>
            </Grid>
          </Container>
        </main>
      </div>
    );
  }

  loadGraphData() {
    // console.log("run axios get on: /measures")
    axios.get('/measures', {
      headers: {
        'Access-Control-Allow-Origin': '*'
      }
    }).then(response => {

      // SORT DATA ASCENDING CONSIDERING TIME 
      let sortedData = response.data.sort((function (b, a) {
        return new Date(b.timestamp) - new Date(a.timestamp)
      }));
      this.setState({ measures: sortedData })
    }, error => {
      console.log(error);
    });
  }
}

export default withStyles(styles)(App);