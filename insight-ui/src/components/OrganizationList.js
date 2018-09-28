import React, { Component } from 'react';

import { NavLink, withRouter } from 'react-router-dom'
import { MenuList, MenuItem, ListItemIcon,
    ListItemText, Divider, ListSubheader
} from '@material-ui/core';
import { Home } from '@material-ui/icons';

class OrganizationList extends Component {
    getMenuItem(item) {
        const url = `/organization/${item.name}`
        return (
            <MenuItem
                key={item.id}
                component={NavLink}
                to={url}
                selected={url === this.props.location.pathname}
            >
                <ListItemText primary={item.name} />
            </MenuItem>
        )
    }

    getHomeItem() {
        const url = "/"
        return (
            <MenuItem
                key="home"
                component={NavLink}
                to={url}
                selected={url === this.props.location.pathname}>
                <ListItemIcon>
                    <Home />
                </ListItemIcon>
                <ListItemText primary="Home" />
            </MenuItem>
        )
    }

    render() {
        return (
            <MenuList>
                { this.getHomeItem() }
                <Divider/>
                <ListSubheader component="div">
                    Organization
                </ListSubheader>
                <Divider/>
                {
                    this.props.organizations.map((item,id) => {
                        if (!item.insight_irma_endpoint || !item.insight_log_endpoint) {
                            return false
                        }

                        return this.getMenuItem({
                            id: id,
                            name: item.name,
                            signed: item.signed
                        });
                    })
                }
            </MenuList>
        );
    }
}

export default withRouter(OrganizationList);