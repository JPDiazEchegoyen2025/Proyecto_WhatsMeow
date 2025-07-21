import React from 'react';

function Loader() {
  return (
    <div style={{ textAlign: 'center', padding: '40px' }}>
      <div className="loader" style={{ margin: '0 auto', width: '48px', height: '48px', border: '6px solid #eee', borderTop: '6px solid #25d366', borderRadius: '50%', animation: 'spin 1s linear infinite' }}></div>
      <p style={{ color: '#afafaf', marginTop: '16px' }}>Cargando...</p>
    </div>
  );
}

export default Loader;
