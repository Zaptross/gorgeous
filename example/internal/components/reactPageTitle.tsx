import react from 'react';
import { useTheme } from '../provider/themeProvider';

type Props = {
  title: string;
};

export function PageTitle({ title }: Props) {
  const { theme } = useTheme();

  return (
    <h1 style={{ color: theme.Base1 }}>
      {title}
    </h1>
  );
}