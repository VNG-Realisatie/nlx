import { css } from 'styled-components'

export const baseButtonStyle = css`
    position: relative;
    height: ${p => p.theme.size.normal};

    background-color: ${p => p.theme.color.white};
    border: 2px solid transparent;
    user-select: none;

    display: inline-flex;
    align-items: center;
    justify-content: center;

    transition: background-color ${p => p.theme.transition.fast}, color ${p => p.theme.transition.fast}, border-color ${p => p.theme.transition.fast};

    ${p => p.variant === 'primary' && css`
        background-color: ${p => p.theme.color.primary.main};
        color: ${p => p.theme.color.white};

        &[disabled] {
            background-color: ${p => p.theme.color.grey[20]};
        }

        &:hover,
        &:focus {
            background-color: ${p => p.theme.color.primary.light};
        }

        &:active {
            background-color: ${p => p.theme.color.primary.dark};
        }
    `}

    ${p => p.variant === 'secondary' && css`
        border-color: ${p => p.theme.color.primary.main};
        color: ${p => p.theme.color.primary.main};

        &[disabled] {
            border-color: ${p => p.theme.color.grey[20]};
        }

        &:hover,
        &:focus {
            background-color: ${p => p.theme.color.primary.lightest};
        }

        &:active {
            background-color: ${p => p.theme.color.primary.lighter};
            border-color: ${p => p.theme.color.primary.main};
        }
    `}

    ${p => p.variant === 'tertiary' && css`
        color: ${p => p.theme.color.black};

        &:hover,
        &:focus {
            background-color: ${p => p.theme.color.grey[10]};
        }

        &:active {
            background-color: ${p => p.theme.color.grey[20]};
        }
    `}

    ${(p) =>
        p.size === 'small' &&
        css`
            height: ${(p) => p.theme.size.small};
    `};

    ${(p) =>
        p.size === 'normal' &&
        css`
            height: ${(p) => p.theme.size.normal};
    `};

    ${(p) =>
        p.size === 'large' &&
        css`
            height: ${(p) => p.theme.size.large};
    `};

    &:not([disabled]) {
        cursor: pointer;
    }

    &[disabled] {
        cursor: default;
        pointer-events: none;
        color: ${p => p.theme.color.grey[50]};
    }
`
