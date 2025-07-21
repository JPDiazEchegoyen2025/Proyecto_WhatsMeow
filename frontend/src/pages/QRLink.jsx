import React, { useEffect } from 'react';
import QRCode from 'react-qr-code';

function QRLink({ qr, qrTimer, onBack, onRefresh }) {
  // Polling para detectar vinculación
  useEffect(() => {
    const interval = setInterval(() => {
      fetch('http://localhost:8080/session')
        .then((res) => res.json())
        .then((data) => {
          if (data && data.active) {
            window.location.reload();
          }
        })
        .catch(() => {});
    }, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <h2 style={{ color: '#202020', marginBottom: '16px' }}>Escanea el QR para iniciar sesión</h2>
        <div style={{ background: '#fff', padding: '24px', display: 'inline-block', borderRadius: '12px', boxShadow: '0 0 8px #eee', marginBottom: '16px' }}>
          <QRCode value={qr} size={300} />
        </div>
        <p style={{ color: '#afafaf', fontSize: '1em', marginBottom: '8px', textAlign: 'center' }}>
          Abre WhatsApp en tu teléfono, ve a <b>Menú &gt; Dispositivos vinculados</b> y escanea el código QR para conectar tu cuenta.
        </p>
        <p style={{ color: '#25d366', fontWeight: 'bold', marginBottom: '8px' }}>El QR expira en: {qrTimer}s</p>
        {qrTimer === 0 && (
          <button onClick={onRefresh} style={{ marginTop: '10px', background: '#25d366', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer' }}>Actualizar QR</button>
        )}
        <button onClick={onBack} style={{ marginTop: '20px' }}>Volver</button>
      </div>
    </div>
  );
}

export default QRLink;
