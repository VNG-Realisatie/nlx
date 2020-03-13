import { node } from 'prop-types'
import styled from 'styled-components'

const Head = styled.thead`
  display: table-header-group;
`

Head.propTypes = {
  children: node,
}

export default Head
