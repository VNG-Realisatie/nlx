import { css } from 'styled-components'

export const buttonStyle = css`
    border-radius: ${p => p.theme.radius.small};

    font-family: ${p => p.theme.font.family.system};
    font-weight: ${p => p.theme.font.weight.semibold};
    letter-spacing: .04em;

    text-transform: uppercase;
    text-shadow: none;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;

     ${(p) =>
        p.size === 'small' &&
        css`
            height: ${(p) => p.theme.size.small};
            padding: 0 0.5rem;
            font-size: ${(p) => p.theme.font.size.tiny};

            > span {
                padding-bottom: 1px;
            }
        `}

    ${(p) =>
        p.size === 'normal' &&
        css`
            height: ${(p) => p.theme.size.normal};
            padding: 0 1rem;
            font-size: ${(p) => p.theme.font.size.small};

            > span {
                padding-bottom: 2px;
            }
        `}

    ${(p) =>
        p.size === 'large' &&
        css`
            height: ${(p) => p.theme.size.large};
            padding: 0 1.5rem;
            font-size: ${(p) => p.theme.font.size.small};

            > span {
                padding-bottom: 2px;
            }
        `}

    ${(p) =>
        p.variant === 'tertiary' &&
        css`
        padding: .5rem;
        `};
`

export const iconStyle = css`
    display: flex;
    margin-right: 6px;

    ${(p) =>
        p.outline !== 'none' &&
        css`
            margin-left: -3px;
        `};

    ${(p) =>
        p.variant === 'secondary' &&
        css`
            color: ${(p) => p.theme.color.primary.main};
        `};

    ${(p) =>
        p.variant === 'tertiary' &&
        css`
            color: ${(p) => p.theme.color.grey[60]};
        `};
`

export const iconRightStyle = css`
    display: flex;
    margin-left: 6px;

    ${(p) =>
        p.outline !== 'none' &&
        css`
            margin-right: -3px;
        `};

    ${(p) =>
        p.variant === 'secondary' &&
        css`
            color: ${(p) => p.theme.color.primary.main};
        `};

    ${(p) =>
        p.variant === 'tertiary' &&
        css`
            color: ${(p) => p.theme.color.grey[60]};
        `};
`
