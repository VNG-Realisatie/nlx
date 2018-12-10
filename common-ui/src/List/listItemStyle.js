import { css } from 'styled-components'

export const listItemStyle = css`
    display: flex;
    align-items: center;

    color: ${(p) => p.theme.color.grey[60]};

    text-decoration: none;
    text-shadow: none;
    user-select: none;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;

    &:hover,
    &:focus {
        background-color: ${(p) => p.theme.color.hover};
    }

    &:active {
        background-color: ${(p) => p.theme.color.active};
    }

    &[aria-current] {
        color: ${(p) => p.theme.color.primary.main};
    }

    > span {
        padding-bottom: ${(p) => p.theme.font.offset.bottom};
    }

    ${(p) =>
        p.size === 'small' &&
        css`
            height: ${(p) => p.theme.size.small};
            padding: 0 1.5rem;

            font-size: ${(p) => p.theme.font.size.small};
            line-height: ${(p) => p.theme.font.lineHeight.small};
        `};

    ${(p) =>
        p.size === 'normal' &&
        css`
            height: ${(p) => p.theme.size.normal};
            padding: 0 1.5rem;

            font-size: ${(p) => p.theme.font.size.normal};
            line-height: ${(p) => p.theme.font.lineHeight.normal};
        `};
`
