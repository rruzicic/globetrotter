
// converts format 'YYYY-mm-DDThh:MM:SS.000Z' to 'DD.mm.YYYY hh:MM'
const formatDate = (dateString) => {
    const date = new Date(dateString);
    const day = date.getDate();
    const month = date.getMonth() + 1;
    const year = date.getFullYear();
    const hours = date.getHours();
    const minutes = date.getMinutes();
    let formattedDate = '';
    if(hours < 10) {
        formattedDate = `${day}.${month}.${year} 0${hours}:${minutes}0`;
    } else {
        formattedDate = `${day}.${month}.${year} ${hours}:${minutes}0`;
    }
    return formattedDate;
}

export default formatDate