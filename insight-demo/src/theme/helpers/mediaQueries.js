import { css } from 'styled-components'

import themeConstants from '../themeConstants'

const { breakpoints, unit, step } = themeConstants.breakpoints

const keys = Object.keys(breakpoints)
const mediaQueries = {}

// up(), down() and only() inspired by material-ui
// https://github.com/mui-org/material-ui/blob/master/packages/material-ui/src/styles/createBreakpoints.js

export const up = (key) => `@media (min-width: ${breakpoints[key]}${unit})`

export const down = (key) => {
    const nextKey = keys[keys.indexOf(key) + 1]

    // down('xl') equals up('xs')
    if (!nextKey) return up(keys[0])

    return `@media (max-width: ${breakpoints[nextKey] - step / 100}${unit})`
}

export const only = (key) => {
    const nextKey = keys[keys.indexOf(key) + 1]

    // only('xl') equals up('xl')
    if (!nextKey) return up(key)

    return (
        `@media (min-width: ${breakpoints[key]}${unit}) and ` +
        `(max-width: ${breakpoints[nextKey] - step / 100}${unit})`
    )
}

const queryFactory = (args, query) => css`
    ${query} {
        ${css(...args)}
    }
`

for (let i = 0, l = keys.length; i < l; i++) {
    const label = keys[i]
    const breakpoint = breakpoints[label]

    if (typeof breakpoint !== 'number') {
        throw Error('Theme breakpoint is not a number')
    }

    // Even though xsDown and xlUp make no sense, just create them so there's no error when they're called
    // Functions above will make sure the styling goes well
    mediaQueries[`${label}Down`] = (...args) => queryFactory(args, down(label))
    mediaQueries[`${label}`] = (...args) => queryFactory(args, only(label))
    mediaQueries[`${label}Up`] = (...args) => queryFactory(args, up(label))
}

/**
 * Use like this:
 *
 * import styled from 'styled-components'
 * import { media } from 'theme/helpers'
 * const MediaTest = styled.div`
 *     background-color: hotpink;
 *
 *     ${media.sm`
 *         background-color: lime;
 *     `}
 *
 *     ${media.lgUp`
 *         background-color: cyan;
 *     `}
 * `
 */
export default mediaQueries
