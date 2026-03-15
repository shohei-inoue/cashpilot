'use client';

import { useState } from 'react';
import styles from './MainContainer.module.scss';
import Header from '../Header/Header';
import Sidebar from '../Sidebar/Sidebar';

const MainContainer = ({ children }: { children: React.ReactNode }) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true);

  return (
    <div className={styles.container}>
      <Header
        isSidebarOpen={isSidebarOpen}
        onToggleSidebar={() => setIsSidebarOpen((prev) => !prev)}
      />
      <div className={styles.body}>
        <Sidebar isOpen={isSidebarOpen} />
        <main className={styles.main}>{children}</main>
      </div>
    </div>
  );
};

export default MainContainer;
