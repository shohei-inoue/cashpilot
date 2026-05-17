'use client';

import { useState } from 'react';
import styles from './MainContainer.module.scss';
import Header from '../Header/Header';
import Sidebar from '../Sidebar/Sidebar';

type MainContainerProps = {
  children: React.ReactNode;
  userEmail: string;
};

const MainContainer = ({ children, userEmail }: MainContainerProps) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true);

  return (
    <div className={styles.container}>
      <Header
        isSidebarOpen={isSidebarOpen}
        onToggleSidebar={() => setIsSidebarOpen((prev) => !prev)}
        userEmail={userEmail}
      />
      <div className={styles.body}>
        <Sidebar isOpen={isSidebarOpen} />
        <main className={styles.main}>{children}</main>
      </div>
    </div>
  );
};

export default MainContainer;
