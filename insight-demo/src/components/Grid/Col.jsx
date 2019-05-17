// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { PureComponent } from 'react'
import PropTypes from 'prop-types'
import { Box } from '@rebass/grid'

/**
 * Usage:
 * Column will be half screen width on small screens and quarter width for medium screens and larger
 * <Col sm={1 / 2} mdUp={1 / 4}>content</Col>
 */
class Col extends PureComponent {
    render() {
        const { children, ...sizes } = this.props

        // Default to full width for all screen sizes
        const width = [
            1 / 1, // xs
            1 / 1, // sm
            1 / 1, // md
            1 / 1, // lg
        ]

        // Define for each possible size prop, what position(s) in `width` array should be changed
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
                // Create new array `values` with passed size (eg. 1 / 2)
                const values = Array(times).fill(sizes[prop])
                // Replace the default width values with the passed value
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
    px: '16px',
    xsUp: undefined,
    smUp: undefined,
    mdUp: undefined,
    xs: undefined,
    sm: undefined,
    md: undefined,
    lg: undefined,
}

export default Col
