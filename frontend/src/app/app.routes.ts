import { Routes } from '@angular/router';

export const routes: Routes = [
    {
        path: '',
        loadComponent: () => import('./components/search/search').then(m => m.SearchComponent)
    },
    { path: '**', redirectTo: '' }
];
