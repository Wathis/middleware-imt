/* eslint-disable no-script-url */

import React from 'react';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import 'typeface-roboto';
import Title from './Title';
import axios from 'axios';


const styles = theme => ({
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
});

class Moyennes extends React.Component {

  constructor() {
    super();
    this.state = {
      averages: {},
      wind: {
        isLoaded: false,
        average: null
      },
      temperature: {
        isLoaded: false,
        average: null
      },
      pressure: {
        isLoaded: false,
        average: null
      },
    }
  }

  componentDidMount() {
    let t = Math.round(new Date() / 1000)
    t = "1571858366"
    console.log("run axios get on: /measures/" + t + "/average")
    axios.get('/measures/' + t + '/average', {
      headers: {
        'Access-Control-Allow-Origin': '*'
      }
    })
      .then(response => {
        console.log(response.data);
        this.setState({ averages: response.data.averages })
        if (response.data.averages.temperature) {
          this.setState({ temperature: { isLoaded: true, average: response.data.averages.temperature } })
        }
        if (response.data.averages.wind) {
          this.setState({ wind: { isLoaded: true, average: response.data.averages.wind } })
        }
        if (response.data.averages.pressure) {
          this.setState({ pressure: { isLoaded: true, average: response.data.averages.pressure } })
        }
      }, error => {
        console.log(error);
      });
  }

  render() {
    const { classes } = this.props;

    let displayMeasures = {
      "temperature": "Temperature : ",
      "pressure": "Pression : ",
      "wind": "Vent : ",
    }
    this.state.temperature.isLoaded ? displayMeasures.temperature += this.state.temperature.average + " °c" : displayMeasures.temperature += 'Pas de données'
    this.state.pressure.isLoaded ? displayMeasures.pressure += this.state.pressure.average + " Pa" : displayMeasures.pressure += 'Pas de données'
    this.state.wind.isLoaded ? displayMeasures.wind += this.state.wind.average + " km/h" : displayMeasures.wind += 'Pas de données'

    return (
      <>
        <Title>Moyennes des relevés aujourd'hui</Title>
        <Grid container spacing={3}>
          <Grid item xs={4}>
            <Paper className={classes.paper}>
              <Typography variant="h6">{displayMeasures.temperature}</Typography>
            </Paper>
          </Grid>
          <Grid item xs={4}>
            <Paper className={classes.paper}>
              <Typography variant="h6">{displayMeasures.wind}</Typography>
            </Paper>
          </Grid>
          <Grid item xs={4}>
            <Paper className={classes.paper}>
              <Typography variant="h6">{displayMeasures.pressure}</Typography>
            </Paper>
          </Grid>
        </Grid>
      </>
    )
  }
}

export default withStyles(styles)(Moyennes);