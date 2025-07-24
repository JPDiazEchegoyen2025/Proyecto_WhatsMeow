


import styles from './ContactsPanel.module.css';
import React, { useState } from 'react';

function ContactItem({ contact, active, onSelect }) {
  const getAvatarColor = (name) => {
    const colors = [
      '#ff6b6b', '#4ecdc4', '#45b7d1', '#96ceb4',
      '#ffeaa7', '#dda0dd', '#98d8c8', '#f7dc6f'
    ];
    const index = name.charCodeAt(0) % colors.length;
    return colors[index];
  };
  const getInitials = (name) => {
    if (!name) return '';
    // Extraer solo letras alfabéticas
    const letters = name.match(/[a-zA-ZáéíóúüñÁÉÍÓÚÜÑ]/g);
    if (!letters || letters.length === 0) return '';
    // Tomar la primera y segunda letra alfabética si existen
    const initials = (letters[0] || '') + (letters[1] || '');
    return initials.toUpperCase();
  };
  return (
    <div
      className={`${styles['contact-item']}${active ? ' ' + styles['active'] : ''}`}
      onClick={() => onSelect(contact)}
      key={contact.id}
    >
      <div
        className={styles['contact-avatar']}
        style={{ backgroundColor: getAvatarColor(contact.name) }}
      >
        {getInitials(contact.name)}
      </div>
      <div className={styles['contact-info']}>
        <div className={styles['contact-header']}>
          <span className={styles['contact-name']}>
            {contact.name}
          </span>
          {contact.lastMessageTime && (
            <span className={styles['message-time']}>
              {contact.lastMessageTime}
            </span>
          )}
        </div>
        <div className={styles['contact-footer']}>
          <span className={styles['last-message']}>
            {contact.lastMessage}
          </span>
          {contact.unreadCount > 0 && (
            <div className={styles['unread-count']}>
              {contact.unreadCount}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
const getAvatarColor = (name) => {
  const colors = [
    '#ff6b6b', '#4ecdc4', '#45b7d1', '#96ceb4',
    '#ffeaa7', '#dda0dd', '#98d8c8', '#f7dc6f'
  ];
  const index = name.charCodeAt(0) % colors.length;
  return colors[index];
};

const getInitials = (name) => {
  return name
    .split(' ')
    .map(n => n[0])
    .join('')
    .substring(0, 2)
    .toUpperCase();
};

export default function ContactsPanel({ contacts, onSelectContact, activeContactId }) {
  const [searchValue, setSearchValue] = useState('');
  const safeContacts = contacts || [];
  const filteredContacts = safeContacts.filter(contact =>
    contact.name && contact.name.toLowerCase().includes(searchValue.toLowerCase())
  );

  return (
    <div className={styles['contacts-panel']}>
      {/* Buscador */}
      <div className={styles['search-container']}>
        <div className={styles['search-box']}>
          <span className={styles['search-icon']}>🔍</span>
          <input
            type="text"
            className={styles['search-input']}
            placeholder="Buscar en chat"
            value={searchValue}
            onChange={e => setSearchValue(e.target.value)}
          />
        </div>
      </div>

      {/* Lista de contactos */}
      <div className={styles['contacts-list']}>
        {filteredContacts.map(contact => (
          <ContactItem
            key={contact.id}
            contact={contact}
            active={activeContactId === contact.id}
            onSelect={onSelectContact}
          />
        ))}
      </div>
    </div>
  );
}
