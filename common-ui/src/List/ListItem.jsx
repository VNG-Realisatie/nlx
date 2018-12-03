import React, { PureComponent } from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

import { listItemStyle } from './listItemStyle.js'

const StyledListItem = styled.div`
    ${listItemStyle};
`

class ListItem extends PureComponent {
    render() {
        return (
            <StyledListItem {...this.props}>
                <span>{this.props.children}</span>
            </StyledListItem>
        )
    }
}

ListItem.propTypes = {
    size: PropTypes.oneOf(['small', 'normal']),
}

ListItem.defaultProps = {
    size: 'small'
}

export default ListItem
