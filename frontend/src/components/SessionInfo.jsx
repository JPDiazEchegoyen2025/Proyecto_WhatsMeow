import React from 'react';

function SessionInfo({ sessionInfo, onLogout }) {
  return (
    <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
      <h2 style={{ color: '#202020', marginBottom: '16px' }}>¡Ya tienes una sesión activa!</h2>
      <p style={{ color: '#202020', fontSize: '1.1em', marginBottom: '16px' }}>
        <b>Usuario:</b> {sessionInfo.user || 'Desconocido'}<br />
        <b>Teléfono:</b> {sessionInfo.phone || 'Desconocido'}<br />
        <b>Estado:</b> {sessionInfo.status || 'Activa'}
      </p>
      <button onClick={onLogout} style={{ marginTop: '20px', background: '#d32f2f', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer' }}>Cerrar sesión</button>
    </div>
  );
}

export default SessionInfo;
