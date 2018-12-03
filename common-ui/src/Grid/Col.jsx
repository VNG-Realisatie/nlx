import React, { PureComponent } from 'react'
import PropTypes from 'prop-types'
import { Box } from '@rebass/grid'

class Col extends PureComponent {
    render() {
        const { children, ...sizes } = this.props

        const width = [
            1 / 1, // xs
            1 / 1, // sm
            1 / 1, // md
            1 / 1, // lg
        ]

        const inserts = new Map([
            ['xsUp', [0, 4]],
            ['smUp', [1, 3]],
            ['mdUp', [2, 2]],
            ['xs', [0, 1]],
            ['sm', [1, 1]],
            ['md', [2, 1]],
            ['lg', [3, 1]],
        ])

        inserts.forEach((insert, prop) => {
            if (sizes[prop]) {
                const [at, times] = insert
                const values = Array(times).fill(sizes[prop])
                width.splice(at, times, ...values)
            }
        })

        return (
            <Box {...this.props} width={width} mb={2}>
                {children}
            </Box>
        )
    }
}

Col.propTypes = {
    children: PropTypes.node.isRequired,
    xsUp: PropTypes.number,
    smUp: PropTypes.number,
    mdUp: PropTypes.number,
    xs: PropTypes.number,
    sm: PropTypes.number,
    md: PropTypes.number,
    lg: PropTypes.number,
}

Col.defaultProps = {
    px: '10px',
    xsUp: undefined,
    smUp: undefined,
    mdUp: undefined,
    xs: undefined,
    sm: undefined,
    md: undefined,
    lg: undefined,
}

export default Col
