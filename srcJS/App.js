import logo from './logo.svg';
import {BrowserRouter as Router,Route, Switch} from 'react-router-dom';
import Rendered from "./Rendered";
import './App.css';
import Form from "./Form";

function App() {
  return (
    <div className="App" >
        <Router>
            <Switch>
                <Route path={'/'} exact={true} component={Form}/>
                <Route path={'/image/:name'} component={Rendered}/>
            </Switch>
        </Router>
    </div>
  );
}

export default App;
