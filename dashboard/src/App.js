// App.js
import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import LoginScreen from './screens/LoginScreen';

const App = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={LoginScreen} />
        {/* Add more routes for your screens */}
      </Switch>
    </Router>
  );
};

export default App;