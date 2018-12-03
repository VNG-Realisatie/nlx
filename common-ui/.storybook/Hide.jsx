import styled from 'styled-components'
import theme from './theme'

const { breakpoints } = theme
const lastIndex = breakpoints.length - 1

const getMediaBreakpoint = (breakpoints, breakpoint, index) => {
  if (index === 0) return `@media screen and (max-width: ${breakpoint})`
  const prevBreakpoint = breakpoints[index - 1]
  if (index === lastIndex) {
    return `@media screen and (min-width: ${prevBreakpoint})`
  }
  return `@media screen and (min-width: ${prevBreakpoint}) and (max-width: ${breakpoint})`
}

const mediaBreakpoints = breakpoints.reduce((acc, breakpoint, index) => {
  acc[index] = getMediaBreakpoint(breakpoints, breakpoint, index)
  return acc
}, {})

const hidden = key => props => {
  const breakpoints = [].concat(props.breakpoints)
  return breakpoints.includes(key)
    ? {
      [mediaBreakpoints[key]]: {
        display: 'none'
      }
    }
    : null
}

const Hide = styled.div(
  [],
  ...Object.keys(mediaBreakpoints).map(i => hidden(Number(i)))
)

export default Hide