import React, { PureComponent } from 'react'
import { Flex } from '@rebass/grid'

class Row extends PureComponent {
    render() {
        const { children } = this.props

        return <Flex {...this.props}>{children}</Flex>
    }
}

Row.propTypes = {}

Row.defaultProps = {
    mx: '-10px',
    flexWrap: 'wrap',
}

export default Row
