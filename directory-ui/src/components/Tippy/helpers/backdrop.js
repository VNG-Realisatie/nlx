import { css } from 'styled-components'

export function backdropTransformEnter(placement) {
    const scale = 1

    let translate
    switch (placement) {
        case 'top':
            translate = '-50%, -55%'
            break
        case 'bottom':
            translate = '-50%, -45%'
            break
        case 'left':
            translate = '-50%, -50%'
            break
        case 'right':
            translate = '-50%, -50%'
            break
        default:
            // no effect
    }

    return css`
        transform: scale(${scale}) translate(${translate});
    `
}

export function backdropTransformLeave(placement) {
    const scale = 0.2

    let translate
    switch (placement) {
        case 'top':
            translate = '-50%, -45%'
            break
        case 'bottom':
            translate = '-50%, 0'
            break
        case 'left':
            translate = '-75%, -50%'
            break
        case 'right':
            translate = '-25%, -50%'
            break
        default:
            // no effect
    }

    return css`
        transform: scale(${scale}) translate(${translate});
    `
}
