import React from 'react';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import { withStyles } from '@material-ui/core/styles';
import 'typeface-roboto';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Grid from '@material-ui/core/Grid';

const styles = theme => ({
    paper: {
        padding: theme.spacing(2),
        textAlign: 'center',
        color: theme.palette.text.secondary,
    },
    spacingTop: {
        marginTop: "50px",
    }
});

class SelectAirport extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            airportId: props.airportList[0],
        }
        this.handleChange = this.handleChange.bind(this);
    }

    handleChange(e) {
        this.setState({ airportId: e.target.value })
        this.props.onAirportChange(e.target.value);
    }

    render() {
        const { classes } = this.props;
        var airportList = this.props.airportList

        return (
            <>
                <Grid container className={classes.spacingTop} spacing={3}>
                    <Grid item xs={12}>
                        <Paper className={classes.paper}>
                            <Typography variant="h6">Selectionner un aéroport</Typography>
                            <InputLabel htmlFor="airport-simple"></InputLabel>
                            <Select
                                value={this.state.airportId}
                                onChange={this.handleChange}
                                className={classes.centerText}
                            >
                                {airportList.map((val, i) => {
                                    return (<MenuItem key={"select-menuitem-" + i} value={val}>{val}</MenuItem>)
                                })}
                            </Select>
                        </Paper>
                    </Grid>
                </Grid>
            </>
        )
    }
}

export default withStyles(styles)(SelectAirport);