import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { TableSettingsProvider } from './context/TableSettingsContext';
import Home from './pages/Home/Home';
import JobsDashboard from './pages/JobsDashboard/JobsDashboard';
import NodesDashboard from './pages/NodesDashboard/NodesDashboard';
import Settings from './pages/Settings/Settings';

const App = () => {
  return (
    <TableSettingsProvider>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/JobsDashboard" element={<JobsDashboard />} />
          <Route path="/NodesDashboard" element={<NodesDashboard />} />
          <Route path="/Settings" element={<Settings />} />
        </Routes>
      </Router>
    </TableSettingsProvider>
  );
};

export default App;