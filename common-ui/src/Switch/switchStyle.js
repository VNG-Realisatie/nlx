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

        &:focus + label {
            &:after {
                border-color: ${(p) => p.theme.color.grey[50]};
            }
        }

        &:checked + label {
            &:before {
                background-color: ${(p) => p.theme.color.primary.main};
            }
            &:after {
                border-color: ${(p) => p.theme.color.primary.main};
                transform: translateX(15px);
            }
        }
    }

    &[disabled] + label {
        color: ${(p) => p.theme.color.grey[50]};
    }
`

export const labelStyle = css`
    padding: 0 0 0 46px;

    font-size: ${(p) => p.theme.font.size.normal};
    line-height: ${(p) => p.theme.font.lineHeight.normal};

    user-select: none;
    pointer-events: none;

    &:before {
        content: '';
        position: absolute;
        display: block;
        width: 34px;
        height: 18px;
        top: 4px;
        left: 0;
        border: 0;
        border-radius: 9px;
        background-color: ${(p) => p.theme.color.grey[30]};
        transition: background-color ${(p) => p.theme.transition.materialNormal};
    }

    &:after {
        content: '';
        position: absolute;
        display: block;
        width: 20px;
        height: 20px;
        top: 3px;
        left: 0;
        border-radius: 50%;
        background-color: white;
        border: 2px solid ${(p) => p.theme.color.grey[40]};
        transform: translateX(-1px);
        transition: transform ${(p) => p.theme.transition.materialNormal};
    }
`

