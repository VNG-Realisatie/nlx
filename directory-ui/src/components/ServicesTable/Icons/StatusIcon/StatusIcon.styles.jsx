import styled from 'styled-components'

const online = '#63D19E'
const offline = '#CAD0E0'

export const StyledSvg = styled.svg`
  display: inline-block;
  width: 20px;
  height: 20px;
`

export const StyledCircle = styled.circle`
  stroke: ${ p => p.status === 'online' ? online : offline}
`
