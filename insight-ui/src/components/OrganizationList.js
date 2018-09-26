import React, { Component } from 'react';
import axios from 'axios';

import { NavLink, withRouter } from 'react-router-dom'
import { MenuList, MenuItem, ListItemIcon, 
    ListItemText, Divider, ListSubheader
} from '@material-ui/core';
import { Home, VerifiedUser, NotInterested }from '@material-ui/icons';

import config from '../utils/config';
import listServices from '../mockdata/directory.dev.nlx-list-services.json';

class OrganizationList extends Component {    
    state={
        "services": []
    }    
    componentDidMount(){                
        //get list of avaliable services        
        this.getServices();
        /*
        this.setState({
            services: listServices.services
        })*/
    }  
    
    getServices(){
        let url = config.api.listServices();
        axios.get(url)
        .then(d =>{
            //debugger            
            this.setState({
                services:d.data.services
            });
        },(e)=>{            
            this.setState({
                "services": []
            });
            console.error(e)
        })
    }
    getMenuItem(item){
        const uri = `/organization/${item.id}`,
            active = uri===this.props.location.pathname;           
        return (
            <MenuItem
                key={item.id}
                component={NavLink} 
                to={uri}
                selected={active}
                title={item.signed ? "Signed in to service" : "Not signed to service"}>
                <ListItemIcon>
                    {item.signed ? (
                        <VerifiedUser />
                    ):(
                        <NotInterested/>
                    )}                    
                </ListItemIcon>
                <ListItemText primary={item.name} />
            </MenuItem> 
        )
    }
    getHomeItem(){
        let active = "/" === this.props.location.pathname;   
        return (
            <MenuItem
                key="home"
                component={NavLink} 
                to="/"
                selected={active}>
                <ListItemIcon>
                    <Home />
                </ListItemIcon>
                <ListItemText primary="Home" />
            </MenuItem> 
        )
    }

    render() {
        let { services } = this.state;
        //console.log("services...", services);
        return (         
            <MenuList>
                { this.getHomeItem() }
                <Divider/>
                <ListSubheader component="div">
                    Services
                </ListSubheader>
                <Divider/>
                { 
                    services.map((item,id)=>{
                        return this.getMenuItem({
                            id: id,
                            name: item.organization_name,
                            signed: item.signed   
                        });
                    })    
                }                
            </MenuList> 
        );
    }
}

//export default withRouter(MenuList);
export default withRouter(OrganizationList);