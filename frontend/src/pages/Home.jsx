import React from 'react';

function Home({ sessionInfo, onLogout, onStart, loading, error }) {
  if (loading) {
    return <div style={{ color: '#afafaf', marginTop: '40px' }}>Verificando sesión...</div>;
  }
  if (error) {
    return <div style={{ color: '#d32f2f', marginTop: '40px' }}>Error: {error}</div>;
  }
  return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      {sessionInfo ? (
        <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
          <h2 style={{ color: '#202020', marginBottom: '16px' }}>¡Ya tienes una sesión activa!</h2>
          <p style={{ color: '#202020', fontSize: '1.1em', marginBottom: '16px' }}>
            <b>Usuario:</b> {sessionInfo.user || 'Desconocido'}<br />
            <b>Teléfono:</b> {sessionInfo.phone || 'Desconocido'}<br />
            <b>Estado:</b> {sessionInfo.status || 'Activa'}
          </p>
          <button onClick={onLogout} style={{ marginTop: '20px', background: '#d32f2f', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer' }}>Cerrar sesión</button>
        </div>
      ) : (
        <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
          <h2 style={{ color: '#202020', marginBottom: '16px' }}>¡Bienvenido a WhatsApp Usando WhatsMeow!</h2>
          <p style={{ color: '#afafaf', fontSize: '1em', maxWidth: '500px' }}>
            WhatsMeow es una librería open source para interactuar con WhatsApp Web desde aplicaciones personalizadas, permitiendo automatizar, enviar y recibir mensajes, y gestionar sesiones de forma segura y flexible.
          </p>
          <button onClick={onStart} className="start-btn" style={{ marginTop: '20px', background: '#25d366', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer' }}>Comenzar</button>
        </div>
      )}
    </div>
  );
}

export default Home;
