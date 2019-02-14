import { css } from 'styled-components'

export const inputStyle = css`
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    margin: 0;
    opacity: 0;

    &[disabled] + label {
        color: #B4B4B4;
    }

    &:not([disabled]) {
        cursor: pointer;

        &:focus + label {
            &:after {
                border-color: #B4B4B4;
            }
        }

        &:checked + label {
            &:before {
                background-color: #517FFF;
            }
            &:after {
                width: 18px;
                height: 18px;
                top: -1px;
                border-color: #517FFF;
                transform: translateX(15px);
            }
        }
    }
`

export const labelStyle = css`
    padding: 0 0 0 46px;

    font-size: 1rem;
    line-height: 1.5rem;

    user-select: none;
    pointer-events: none;

    &:before {
        content: '';
        position: absolute;
        display: block;
        width: 34px;
        height: 18px;
        top: 1px;
        left: 0;
        border: 0;
        border-radius: 9px;
        background-color: #EAEAEA;
        transition: background-color 0.25s ease;
    }

    &:after {
        content: '';
        position: absolute;
        display: block;
        width: 16px;
        height: 16px;
        top: 0px;
        left: 0;
        border-radius: 50%;
        background-color: white;
        border: 2px solid #DADADA;
        transform: translateX(-1px);
        transition: transform 0.25s ease;
    }
`
