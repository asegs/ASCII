export default function calculate(latitude,zoom){
    let circumference = Math.abs(Math.cos(latitude))*12756*1000*Math.PI;
    let range = 71*(Math.pow(2,21-zoom));
    return range / circumference * 180;
}
