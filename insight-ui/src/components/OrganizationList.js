// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'

import { NavLink } from 'react-router-dom'
import {
    MenuList,
    MenuItem,
    ListItemIcon,
    ListItemText,
    Divider,
    ListSubheader,
    Button,
} from '@material-ui/core'
import { Home } from '@material-ui/icons'

class OrganizationList extends Component {
    pushToLocation = (item) => {
        let url = `/organization/${item.name}`
        this.props.history.push(url)
    }

    getMenuButton(item) {
        return (
            <Button
                color="primary"
                key={item.id}
                onClick={() => {
                    this.pushToLocation(item)
                }}
            >
                {item.name}
            </Button>
        )
    }
    getMenuItem(item) {
        const url = `/organization/${item.name}`
        return (
            <MenuItem
                key={item.id}
                component={NavLink}
                to={url}
                selected={url === window.location.href}
            >
                <ListItemText primary={item.name} />
            </MenuItem>
        )
    }

    getHomeItem() {
        const url = '/'
        return (
            <MenuItem
                key="home"
                component={NavLink}
                to={url}
                selected={url === window.location.href}
            >
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
                {this.getHomeItem()}
                <Divider />
                <ListSubheader component="div">Organization</ListSubheader>
                <Divider />
                {this.props.organizations.map((item, id) => {
                    return this.getMenuItem({
                        id: id,
                        name: item.name,
                        signed: item.signed,
                        organization: item,
                    })
                })}
            </MenuList>
        )
    }
}

export default OrganizationList
