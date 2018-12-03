import { css } from 'styled-components'

export const inputStyle = css`
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    margin: 0;
    opacity: 0;

    &:not([disabled]) {
        cursor: pointer;

        &:focus + label + div {
            background-color: ${(p) => p.theme.color.hover};
            border-color: ${(p) => p.theme.color.grey[50]};
        }

        &:checked + label + div {
            color: ${(p) => p.theme.color.primary.main};
            border-color: ${(p) => p.theme.color.primary.main};
        }

        &:checked:focus + label + div {
            background-color: ${(p) => p.theme.color.primary.lightest};
        }
    }

    &[disabled] + label {
        color: ${(p) => p.theme.color.grey[50]};
    }
`

export const boxStyle = css`
    display: flex;
    justify-content: center;
    align-items: center;

    position: absolute;
    top: 4px;
    left: 0;

    width: 18px;
    height: 18px;
    color: transparent;

    border: 2px solid ${(p) => p.theme.color.grey[40]};
    border-radius: ${(p) => p.theme.radius.small};
    pointer-events: none;
`

export const labelStyle = css`
    display: block;
    padding: 0 0 0 28px;

    font-size: ${(p) => p.theme.font.size.normal};
    line-height: ${(p) => p.theme.font.lineHeight.normal};
    color: ${(p) => p.theme.color.black};
`
