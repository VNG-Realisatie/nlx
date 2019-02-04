import styled from 'styled-components'

const green = 'rgb(179, 232, 123)'
const red = 'rgb(255, 130, 130)'

const StyledSvg = styled.svg`
  display: inline-block;
  width: 10px;
  height: 10px;
  color: ${ p => p.status === 'online' ? green : red}
`

export default StyledSvg
