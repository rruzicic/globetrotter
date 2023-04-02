
export const formatLocaleDate = (dateString) => {
    let date = new Date(dateString)
    date.setUTCMinutes(date.getUTCMinutes() - date.getTimezoneOffset())
    date.setUTCSeconds(0)
    return `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`
}
