import React,{Component} from "react";

class PhotoElement extends Component{
    constructor(props) {
        super(props);
        this.state = {

        }

    }
    render() {
        let newUrl = window.location.href+"image/"+this.props.name;
        return(
            <div>
                Photo taken at: <button onClick={()=>this.props.jumpTo(this.props.latitude,this.props.longitude)}>{this.props.latitude},{this.props.longitude}</button> <button onClick={()=>window.open(newUrl, '_blank')}>View image</button>
            </div>
        )
    }
}
export default PhotoElement;
