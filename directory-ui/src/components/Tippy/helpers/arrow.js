import { css } from 'styled-components'

export function arrow(placement, width, color) {
    if (placement === 'top') {
        return css`
            border-top: ${width} solid ${color};
            border-right: ${width} solid transparent;
            border-left: ${width} solid transparent;
        `
    } else if (placement === 'bottom') {
        return css`
            border-bottom: ${width} solid ${color};
            border-right: ${width} solid transparent;
            border-left: ${width} solid transparent;
        `
    } else if (placement === 'left') {
        return css`
            border-left: ${width} solid ${color};
            border-top: ${width} solid transparent;
            border-bottom: ${width} solid transparent;
        `
    } else if (placement === 'right') {
        return css`
            border-right: ${width} solid ${color};
            border-top: ${width} solid transparent;
            border-bottom: ${width} solid transparent;
        `
    }
}

export function arrowPosition(placement, width, isRoundArrow) {
    if (placement === 'top') {
        return css`
            bottom: -${width};
        `
    } else if (placement === 'bottom') {
        return css`
            top: -${width};
        `
    }

    if (isRoundArrow) {
        if (placement === 'left') {
            return css`
                right: -${width} * 2;
            `
        } else if (placement === 'right') {
            return css`
                left: -${width} * 2;
            `
        }
    } else {
        if (placement === 'left') {
            return css`
                right: -${width};
            `
        } else if (placement === 'right') {
            return css`
                left: -${width};
            `
        }
    }
}

export function arrowMargin(placement) {
    if (placement === 'top' || placement === 'bottom') {
        return css`
            margin: 0 6px;
        `
    } else {
        return css`
            margin: 3px 0;
        `
    }
}

export function roundarrowTransform(placement) {
    if (placement === 'top') {
        return css`
            transform: rotate(180deg);
        `
    } else if (placement === 'bottom') {
        return css`
            transform: rotate(0);
        `
    } else if (placement === 'left') {
        return css`
            transform: rotate(90deg);
        `
    } else if (placement === 'right') {
        return css`
            transform: rotate(-90deg);
        `
    }
}

export function arrowTransformOrigin(placement) {
    if (placement === 'top') {
        return css`
            transform-origin: 50% 0%;
        `
    } else if (placement === 'bottom') {
        return css`
            transform-origin: 50% 100%;
        `
    } else if (placement === 'left') {
        return css`
            transform-origin: 0% 50%;
        `
    } else if (placement === 'right') {
        return css`
            transform-origin: 100% 50%;
        `
    }
}

export function roundarrowTransformOrigin(placement) {
    if (placement === 'top') {
        return css`
            transform-origin: 50% 0%;
        `
    } else if (placement === 'bottom') {
        return css`
            transform-origin: 50% 100%;
        `
    } else if (placement === 'left') {
        return css`
            transform-origin: 33.33333333% 50%;
        `
    } else if (placement === 'right') {
        return css`
            transform-origin: 66.66666666% 50%;
        `
    }
}
