import styled from 'styled-components'

export const Wrapper = styled.div`
    position: relative;
    display: flex;
`

export const Input = styled.input`
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

        &:checked + label {
            &:before {
                background-color: #517FFF;
            }
            &:after {
                border-color: #517FFF;
                transform: translateX(12px);
            }
        }
    }
`

export const Label = styled.label`
    padding: 0 0 0 38px;

    font-size: 1rem;
    line-height: 1.5rem;

    user-select: none;
    pointer-events: none;

    &:before {
        content: '';
        position: absolute;
        top: 3px;
        left: 0;
        width: 28px;
        height: 16px;
        border: 0;
        border-radius: 8px;
        background-color: #F0F2F7;
        transition: background-color 0.25s ease;
    }

    &:after {
        content: '';
        position: absolute;
        width: 18px;
        height: 18px;
        top: 2px;
        left: -1px;
        border-radius: 50%;
        background-color: white;
        border: 2px solid #CAD0E0;
        transform: translateX(-1px);
        transition: transform 0.25s ease;
    }
`
