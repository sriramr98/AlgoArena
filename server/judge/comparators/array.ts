//TODO: Handle nested arrays and objects
export const areArraysEqualOrdered = (arr1: any[] = [], arr2: any[] = []): boolean => {
    if (arr1.length !== arr2.length) {
        return false;
    }

    if (!Array.isArray(arr1) || !Array.isArray(arr2)) {
        return false;
    }

    for (let i = 0; i < arr1.length; i++) {
        if (arr1[i] !== arr2[i]) {
            return false;
        }
    }

    return true;
}

//TODO: Handle nested arrays and objects
export const areArraysEqualUnordered = (arr1: any[] = [], arr2: any[] = []): boolean => {
    if (arr1.length !== arr2.length) {
        return false;
    }

    if (!Array.isArray(arr1) || !Array.isArray(arr2)) {
        return false;
    }

    const occurences: Record<string, number> = {};
    for (const item of arr1) {
        occurences[item] = (occurences[item] || 0) + 1;
    }
    for (const item of arr2) {
        if (!occurences[item]) {
            return false;
        }
        occurences[item]--;
        if (occurences[item] < 0) {
            return false;
        }
    }

    return true;
}