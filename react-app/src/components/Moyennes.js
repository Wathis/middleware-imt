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
      averages : {}
    }
  }

  componentDidMount() {
    // GET CURRENT TIMESTAMP AND GET MEASURES FROM API
    let t = Math.round(new Date() / 1000)
    // console.log("run axios get on: /measures/" + t + "/average")
    axios.get('/measures/' + t + '/average', {
      headers: {
        'Access-Control-Allow-Origin': '*'
      }
    })
      .then(response => {
        console.log(response.data);
        this.setState(response.data)
      }, error => {
        console.log(error);
      });
  }

  render() {
    const { classes } = this.props;
    var averages = this.state.averages

    let displayMeasures = [
      { sensorType : "WIND", displayMsg : "Vent :", unit : "km.h" },
      { sensorType : "PRESSURE", displayMsg : "Pression :", unit : "Pa" },
      { sensorType : "TEMP", displayMsg : "Température :", unit : "°c" },
    ]

    
    var k
    displayMeasures.forEach(a => {
      k = 0
      Object.keys(averages).forEach(b => {
        if (a.sensorType === b) {
          a.displayMsg += " " + Math.round( averages[b] * 10 ) / 10 + " " + a.unit
          k++
        }
      })
      if (k===0) a.displayMsg += " Pas de données"
    })


    var displayAverages = displayMeasures.map((a,i) => {
      return (
        <Grid item xs={4} key={"grid-average-"+i}>
          <Paper className={classes.paper}>
            <Typography variant="h6">{a.displayMsg}</Typography>
          </Paper>
        </Grid>
      )
    })

    return (
      <>
        <Title>Moyennes des relevés aujourd'hui</Title>
        <Grid container spacing={3}>
          {displayAverages}
        </Grid>
      </>
    )
  }
}

export default withStyles(styles)(Moyennes);