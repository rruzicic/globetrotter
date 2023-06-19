import { Box, Grid, Tooltip, useTheme } from "@mui/material";
import NetworkWifi3BarIcon from '@mui/icons-material/NetworkWifi3Bar';
import PetsIcon from '@mui/icons-material/Pets';
import SoupKitchenIcon from '@mui/icons-material/SoupKitchen';
import LocalParkingIcon from '@mui/icons-material/LocalParking';
import AcUnitIcon from '@mui/icons-material/AcUnit';
import KitchenIcon from '@mui/icons-material/Kitchen';
import TvIcon from '@mui/icons-material/Tv';
import MicrowaveIcon from '@mui/icons-material/Microwave';
import FitnessCenterIcon from '@mui/icons-material/FitnessCenter';

const BenefitsSelectionGrid = ({ selected, setSelected, disabled }) => {
    const theme = useTheme()
    const benefits = [
        { label: 'Wifi', value: 'WIFI', icon: <NetworkWifi3BarIcon fontSize="large" /> },
        { label: 'Gym', value: 'GYM', icon: <FitnessCenterIcon fontSize="large" /> },
        { label: 'Kitchen', value: 'KITCHEN', icon: <KitchenIcon fontSize="large" /> },
        { label: 'Parking', value: 'PARKING', icon: <LocalParkingIcon fontSize="large" /> },
        { label: 'Pets', value: 'PETS', icon: <PetsIcon fontSize="large" /> },
        { label: 'A/C', value: 'AIR_CONDITIONING', icon: <AcUnitIcon fontSize="large" /> },
        { label: 'Micro Wave', value: 'MICROWAVE', icon: <MicrowaveIcon fontSize="large" /> },
        { label: 'TV', value: 'TV', icon: <TvIcon fontSize="large" /> },
        { label: 'Fridge', value: 'FRIDGE', icon: <SoupKitchenIcon fontSize="large" /> },
    ]

    const styles = {
        box: {
            width: "60px",
            height: '60px',
            cursor: 'pointer',
            border: `1px ${theme.palette.primary.main} solid`,
            display: 'grid',
            placeItems: 'center',
            color: theme.palette.primary.main,
            borderRadius: '5px'
        },
        active: {
            backgroundColor: theme.palette.primary.main,
            color: '#fefefe !important'
        }
    }

    const isSelected = (value) => {
        return Boolean(selected.find(benefit => benefit === value));
    }

    const select = (label) => {
        if (disabled) return
        setSelected(prevSelected => {
            const index = prevSelected.indexOf(label);
            if (index === -1) {
                return [...prevSelected, label];
            } else {
                return prevSelected.filter(item => item !== label);
            }
        });
    }

    return (
        <Grid container sx={{ width: '190px', height: '190px' }}>
            {benefits.map((benefit, index) => {
                return (
                    <Grid item xs={4} key={index} sx={{ padding: '0', margin: '0', display: 'grid', placeItems: 'center' }}>
                        <Tooltip title={benefit.label}>
                            <Box onClick={() => select(benefit.label)} sx={isSelected(benefit.label) ? { ...styles.box, ...styles.active } : styles.box}>
                                {benefit.icon}
                            </Box>
                        </Tooltip>
                    </Grid>
                )
            })}
        </Grid>
    );
}

export default BenefitsSelectionGrid;