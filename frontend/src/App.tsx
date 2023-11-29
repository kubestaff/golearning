import './App.css';

import Users from './components/Users';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Users/>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
