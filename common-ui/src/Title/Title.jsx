import PropTypes from 'prop-types'
import styled from 'styled-components'
import {css} from 'styled-components'

const Title = styled.div`
    color: ${p => p.theme.color.black};

    ${p => p.size === 'small' && css`
        font-weight: ${p => p.theme.font.weight.bold};
    `}

    ${p => p.size === 'normal' && css`
        font-weight: ${p => p.theme.font.weight.normal};
        font-size: ${p => p.theme.font.size.title.normal};
        line-height: ${p => p.theme.font.lineHeight.title.normal};
    `}

    ${p => p.size === 'large' && css`
        font-size: ${p => p.theme.font.size.title.large};
        line-height: ${p => p.theme.font.lineHeight.title.large};
    `}
`

Title.propTypes = {
    size: PropTypes.oneOf(['small', 'normal', 'large']),
}

Title.defaultProps = {
    size: 'normal'
}

export default Title