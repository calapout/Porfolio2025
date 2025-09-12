import type {GlobalThemeOverrides} from 'naive-ui'
import colorPalette from "@/ColorPalette.ts";
import ColorPalette from "@/ColorPalette.ts";

function clamp(min: number, max: number, a: number) {
    if (a < min) {
        return min;
    }
    if (a > max) {
        return max;
    }
    return a;
}

function opacityToHex(opacity: number) {
    if (Number.isInteger(opacity)) {
        return clamp(0, 255, opacity).toString(16)
    }

    return Math.round((clamp(0, 1, opacity) * 255)).toString(16)
}


export const ThemeOverrides: GlobalThemeOverrides = {
    common: {
        primaryColor: colorPalette.surface["10"],
        primaryColorHover: colorPalette.surface["40"],
        baseColor: colorPalette.surface["10"],
        textColorBase: colorPalette.text.light,
        textColor1: colorPalette.text.light,
        textColor2: colorPalette.text.light,
        textColor3: colorPalette.text.light,

        borderColor: colorPalette.surface["30"],
        bodyColor: colorPalette.surface["0"],
    },
    Anchor: {
        linkTextColor: colorPalette.text.light,
        linkTextColorHover: colorPalette.text.light,
        linkTextColorActive: colorPalette.text.light,
    },
    Button: {
        color: colorPalette.surface["0"],
    },
    Menu: {
        itemTextColorActiveHorizontal: colorPalette.primary["30"],
        itemTextColorActiveHoverHorizontal: colorPalette.primary["30"],
        itemTextColorHorizontal: colorPalette.text.light,
        itemTextColorHoverHorizontal: colorPalette.text.light,
        itemColorActive: colorPalette.surface["0"],
        itemColorActiveHover: colorPalette.primary["20"] + opacityToHex(0.3),
        itemColorHover: colorPalette.surface["10"],
    },
    Spin:  {
        color: ColorPalette.text.light,
    },
    Card: {
        color: colorPalette.surface["10"],
    },
    Tag: {
        colorPrimary: colorPalette.primary["20"],
        textColorPrimary: colorPalette.text.light,
    },
}