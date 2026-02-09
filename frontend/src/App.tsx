import React, { useState, useEffect } from 'react';
import { IssuesTable } from './components/IssuesTable';
import { WorkloadsTable } from './components/WorkloadsTable';
// Add LatencyTable and SecurityTable components

function App() {
  return (
    <div className="min-h-screen bg-gray-50 p-6">
      <div className="max-w-7xl mx-auto">
        <header className="mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-2">🖥️ Kubewatch</h1>
          <p className="text-xl text-gray-600">Kubernetes Health & Performance Dashboard</p>
        </header>
        
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <div className="bg-white rounded-xl shadow-lg p-8">
            <h2 className="text-2xl font-semibold text-gray-900 mb-6">🚨 Critical Issues</h2>
            <IssuesTable />
          </div>
          
          <div className="bg-white rounded-xl shadow-lg p-8">
            <h2 className="text-2xl font-semibold text-gray-900 mb-6">📊 Workload Metrics</h2>
            <WorkloadsTable />
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
