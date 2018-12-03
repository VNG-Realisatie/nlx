import { css } from 'styled-components'
import { placeholder } from 'polished'

export const inputStyle = css`
    width: 100%;
    height: ${(p) => p.theme.size.large};
    padding: 0 0 ${(p) => p.theme.font.offset.bottom} 1rem;

    background: none;
    color: ${(p) => p.theme.color.black};
    font: inherit;
    border: 2px solid ${(p) => p.theme.color.grey[30]};
    border-radius: ${p => p.theme.radius.small};

    ${placeholder({ color: 'transparent' })};

    &:-webkit-autofill {
        box-shadow: 0 0 0 100px ${(p) => p.theme.color.white} inset;
        -webkit-text-fill-color: ${(p) => p.theme.color.black};
    }

    &:focus {
        border-color: ${(p) => p.theme.color.grey[40]};
    }

    &:not([disabled]) {
        background-color: ${(p) => p.theme.color.white};
    }

    &[disabled] {
        color: ${(p) => p.theme.color.grey[50]};
        background-color: ${(p) => p.theme.color.grey[20]};
        border-color: transparent;
    }
`

export const labelStyle = css`
    display: flex;
    align-items: center;

    position: absolute;
    top: 5px;
    left: 0;
    right: 1px;
    height: ${(p) => p.theme.size.large};
    padding-bottom: ${(p) => p.theme.font.offset.bottom};

    pointer-events: none;

    transition: transform ${(p) => p.theme.transition.materialNormal},
        height ${(p) => p.theme.transition.materialNormal};

    span {
        position: relative;
        margin-left: calc(1rem - 1px);
        padding: 0 2px;

        font-size: ${(p) => p.theme.font.size.normal};
        line-height: 1;
        color: ${(p) => p.theme.color.grey[50]};
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;

        transition: font-size ${(p) => p.theme.transition.materialNormal},
            color ${(p) => p.theme.transition.normal};

        /* White background behind the label */
        &:before {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background-color: ${(p) => p.theme.color.white};
            border-radius: ${(p) => p.theme.radius.small};
            z-index: -1;
            opacity: 0;
            transform: scale(0);
            transform-origin: center;
            transition: opacity ${(p) => p.theme.transition.normal},
                transform ${(p) => p.theme.transition.materialNormal};
        }
    }

    ${(p) =>
        p.small &&
        css`
            height: 20px;
            transform: translateY(-9px);

            span {
                font-size: ${(p) => p.theme.font.size.tiny};

                &:before {
                    opacity: 1;
                    transform: scale(1);
                }
            }
        `};
`
