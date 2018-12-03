import { css } from 'styled-components'

export const iconButtonStyle = css`
    ${p => p.size === 'small' && css`
        width: ${p => p.theme.size.small};
    `}

    ${p => p.size === 'normal' && css`
        width: ${p => p.theme.size.normal};
    `}

    ${p => p.size === 'large' && css`
        width: ${p => p.theme.size.large};
    `}

    border-radius: 50%;
`
