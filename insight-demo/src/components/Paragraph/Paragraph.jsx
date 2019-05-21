// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import styled, { css } from 'styled-components'
import PropTypes from 'prop-types'

const Paragraph = styled.div`
    font-family: ${p => p.theme.font.family.main};

    ${(p) =>
        p.size === 'normal' &&
        css`
            color: ${p => p.theme.color.black};
            font-size: ${p => p.theme.font.size.normal};
            line-height: ${p => p.theme.font.lineHeight.normal};
        `};

    ${(p) =>
        p.size === 'large' &&
        css`
            color: ${p => p.theme.color.grey[60]};
            font-size: ${p => p.theme.font.size.large};
            line-height: ${p => p.theme.font.lineHeight.large};

            span {
                color: ${p => p.theme.color.primary.main};
            }
        `};
`

Paragraph.propTypes = {
    size: PropTypes.oneOf(['normal', 'large']),
}

Paragraph.defaultProps = {
    size: 'normal'
}

export default Paragraph
