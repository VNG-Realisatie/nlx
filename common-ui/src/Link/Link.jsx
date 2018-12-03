import PropTypes from 'prop-types'
import styled from 'styled-components'

import { linkStyle } from './linkStyle.js'

const Link = styled.a`
    ${linkStyle};
`

Link.propTypes = {
    underline: PropTypes.bool,
}

Link.defaultProps = {
    underline: false,
}

export default Link
