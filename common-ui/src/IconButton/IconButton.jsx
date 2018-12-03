import PropTypes from 'prop-types'
import styled from 'styled-components'

import BaseButton from 'src/BaseButton/BaseButton'
import { iconButtonStyle } from './iconButtonStyle.js'

const IconButton = styled(BaseButton)`
    ${iconButtonStyle};
`

IconButton.propTypes = {
    size: PropTypes.oneOf(['small', 'normal', 'large']),
    variant: PropTypes.oneOf(['primary', 'secondary', 'tertiary']),
    icon: PropTypes.element,
    disabled: PropTypes.bool,
}

IconButton.defaultProps = {
    size: 'normal',
    variant: 'primary',
}

export default IconButton