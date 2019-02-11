import styled from 'styled-components'

const grey = '#CAD0E0'
const blue = '#517FFF'

export default styled.svg`
  display: inline-block
  width: 24px
  height: 24px
  fill: ${ p => p.color === 'blue' ? blue : grey };
`