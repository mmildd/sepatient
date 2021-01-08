import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import ApartmentIcon from '@material-ui/icons/Apartment';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import LocalLibraryIcon from '@material-ui/icons/LocalLibrary';
import HealingIcon from '@material-ui/icons/Healing';
import LocalPharmacyIcon from '@material-ui/icons/LocalPharmacy';
import ColorizeIcon from '@material-ui/icons/Colorize';
import AccessibilityIcon from '@material-ui/icons/Accessibility';
import AirlineSeatFlatAngledIcon from '@material-ui/icons/AirlineSeatFlatAngled';
import EmojiPeopleIcon from '@material-ui/icons/EmojiPeople';
import Typography from '@material-ui/core/Typography';
import AirlineSeatIndividualSuiteIcon from '@material-ui/icons/AirlineSeatIndividualSuite';
import {
  Container,
  Grid,
  Table,
} from '@material-ui/core';
import Link from '@material-ui/core/Link';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command

const HeaderCustom = {
  minHeight: '50px',
};
// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 1000,
  },
}));

  
      function a11yProps(index: any) {
        return {
          id: `scrollable-force-tab-${index}`,
          'aria-controls': `scrollable-force-tabpanel-${index}`,
        };
      }

      export default function ScrollableTabsButtonForce() {
        const classes = useStyles();
        const [value, setValue] = React.useState(0);
      
        const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
          setValue(newValue);
        };

  return (

    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`พยาบาล`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
  
        <Link component={RouterLink} to="/">
             Logout
         </Link>
      </Header>

    
    
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={10}>
            <Grid item xs={1}>
   <Table >
     <tr>
       <td  align="center">
           <br></br>
         <ApartmentIcon color="primary" style={{ fontSize: 150 }} />
     <Typography component="h1" variant="h5">
     SE63 Hospital
     </Typography><br></br>
     </td>
     </tr>

     <tr>
         <hr ></hr>
     </tr>
     <Grid item xs={5}>
              <div className={classes.paper}> </div>
            </Grid>
     <tr><br></br></tr>
     <tr>
     <td>
     <AppBar position="static" color="default">
        <Tabs
          value={value}
          onChange={handleChange}
          variant="scrollable"
          scrollButtons="on"
          indicatorColor="primary"
          textColor="primary"
          aria-label="scrollable force tabs example"
        >
           <Tab label="ระบบลงทะเบียนผู้ป่วยใน" icon={<EmojiPeopleIcon />} {...a11yProps(0)}  component={RouterLink} to="/"/>
          <Tab label="ระบบจองห้องพักผู้ป่วย" icon={<AirlineSeatIndividualSuiteIcon />} {...a11yProps(1)} />
          <Tab label="ระบบบันทึกข้อมูลการทำหัตถการให้ผู้ป่วย" icon={<HealingIcon />} {...a11yProps(2)} />
          <Tab label="ระบบการสั่งยา" icon={<LocalPharmacyIcon />} {...a11yProps(3)} />
        </Tabs>
      </AppBar>
     </td>
     </tr>
<br></br><br></br><br></br><br></br>
<tr><Typography component="h1" variant="h5">
     History
     </Typography><br></br></tr>
<br></br>
<tr>
     <td>
     <AppBar position="static" color="default">
        <Tabs
          value={value}
          onChange={handleChange}
          variant="scrollable"
          scrollButtons="on"
          indicatorColor="primary"
          textColor="primary"
          aria-label="scrollable force tabs example"
        >
          <Tab label="ระบบค้นหาข้อมูลการเข้ารักษาผู้ป่วยใน" icon={<AccessibilityIcon />} {...a11yProps(4)}  component={RouterLink} to="/"/>
          <Tab label="ระบบค้นหาห้องพักผู้ป่วย" icon={<AirlineSeatFlatAngledIcon />} {...a11yProps(5)} />
          <Tab label="ระบบค้นหาข้อมูลการทำหัตถการให้ผู้ป่วย" icon={<LocalLibraryIcon />} {...a11yProps(6)} />
          <Tab label="ระบบค้นหาข้อมูลการสั่งยาห้ผู้ป่วยใน" icon={<ColorizeIcon />} {...a11yProps(7)} />
        </Tabs>
      </AppBar>
     </td>
     </tr>
     </Table>
     </Grid> 
           
           
            </Grid>
        </Container>
      </Content>
   
    </Page>
  );
  }