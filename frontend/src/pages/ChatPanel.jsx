import React from 'react';

function ChatPanel({ user, onLogout }) {
  return (
    <header style={{
      width: '100vw',
      position: 'fixed',
      top: 0,
      left: 0,
      zIndex: 100,
      background: '#fff',
      boxShadow: '0 2px 8px #eee',
      minHeight: '60px',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'space-between',
      padding: '0 32px',
      fontWeight: 'bold',
      fontSize: '1.1em',
      borderBottom: '1px solid #e0e0e0'
    }}>
      <span style={{ color: '#1976d2' }}>ChatPanel con WhatsMeow</span>
      <div style={{ display: 'flex', alignItems: 'center', marginLeft: '24px' }}>
        <span style={{ color: '#202020' }}>Usuario: {user || 'Desconocido'}</span>
        <button onClick={onLogout} style={{ background: '#d32f2f', color: '#fff', border: 'none', borderRadius: '8px', padding: '8px 20px', fontWeight: 'bold', cursor: 'pointer', marginLeft: '24px' }}>Cerrar Sesión</button>
      </div>
    </header>
  );
}

export default ChatPanel;
