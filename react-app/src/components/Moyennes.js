/* eslint-disable no-script-url */

import React from 'react';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Container from '@material-ui/core/Container';
import Title from './Title';


class Moyennes extends React.Component {
  constructor() {
    super();
    this.state = {
      wind: {
        isLoaded: false,
        average: null
      },
      temp: {
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
    fetch("http://localhost:8080/measures/wind/average", {method: 'GET', mode: 'no-cors'})
      .then(function(res){
        console.log(res)
        return res
      })
      .then(
        (result) => {
          console.log(result)
        },
        (error) => {
          console.log(error)
        }
      )
  }

  render() {
    return (
      <React.Fragment>
        <Title>Moyenne par type</Title>
        <Container maxWidth="lg" className="">
          <Grid container spacing={3}>
            <Grid item xs={12} md={4} lg={4}>
              <Paper>
                <Typography color="initial">
                  Temperature : {toString(this.state.wind.average)}
                </Typography>
              </Paper>
            </Grid>
            <Grid item xs={12} md={4} lg={4}>
              <Paper>
                <Typography color="initial">
                  Vent :
                </Typography>
              </Paper>
            </Grid>
            <Grid item xs={12} md={4} lg={4}>
              <Paper>
                <Typography color="initial">
                  Pression :
                </Typography>
              </Paper>
            </Grid>
          </Grid>
        </Container>
      </React.Fragment>
    )
  }
}

export default Moyennes;