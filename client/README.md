# Paper Mario 64 Randomizer Statistics - Vue Client

A modern Vue.js client for displaying Paper Mario 64 Randomizer race statistics, built with Vue 3, PrimeVue, TanStack Vue Query, and Tailwind CSS.

## Features

- **Modern Vue 3** with Composition API and TypeScript
- **PrimeVue Components** for rich UI elements (dropdowns, data tables, etc.)
- **TanStack Vue Query** for efficient data fetching and caching
- **Tailwind CSS** with PrimeUI integration for styling
- **Paper Mario Theme** with custom fonts and colors
- **Responsive Design** that works on desktop and mobile

## Prerequisites

- Node.js 20.19+ or 22.12+
- pnpm (recommended) or npm
- Paper Mario Statistics API server running on `http://localhost:3000`

## Installation

```bash
# Install dependencies
pnpm install

# Start development server
pnpm dev

# Build for production
pnpm build

# Preview production build
pnpm preview
```

## Project Structure

```
src/
├── components/          # Vue components
│   ├── StatisticCard.vue    # Individual statistic display
│   └── RawDataTable.vue     # Race data table
├── composables/         # Vue composables for data fetching
│   ├── useStatistics.ts     # Statistics data hooks
│   ├── useEntrants.ts       # Entrants data hooks
│   └── useCategories.ts     # Categories data hooks
├── services/           # API service layer
│   ├── statisticsApi.ts     # Statistics API calls
│   ├── entrantsApi.ts       # Entrants API calls
│   └── categoriesApi.ts     # Categories API calls
├── types.ts            # TypeScript type definitions
├── config.ts           # Application configuration
├── utils.ts            # Utility functions
├── style.css           # Global styles
├── App.vue             # Main application component
└── main.ts             # Application entry point
```

## Key Technologies

### Vue 3 + TypeScript
- Uses Composition API with `<script setup>`
- Full TypeScript support for type safety
- Reactive state management with `ref` and `computed`

### TanStack Vue Query
- Automatic background refetching
- Caching and invalidation
- Loading and error states
- Optimistic updates
- Query key management

### PrimeVue
- **Dropdown**: For filter selections
- **DataTable**: For displaying race data with pagination
- **ProgressSpinner**: For loading indicators
- **Aura Theme**: Modern design system

### API Integration
The client connects to these endpoints:
- `GET /api/get_race_entrants` - Fetch all race participants
- `GET /api/get_statistics_for_entrant` - Fetch statistics with filters
- `GET https://racetime.gg/pm64r/data` - Fetch race categories

## Configuration

Edit `src/config.ts` to modify:
- API base URL
- Default race goal
- Query cache timings
- Race type options

## Filters

The application supports three types of filters:

1. **Racer Filter** - Select a specific racer or "All"
2. **Race Goal Filter** - Filter by race category/goal
3. **Race Type Filter** - Filter by league vs community races

## Statistics Display

Based on selected filters, different statistics are shown:

- **Always**: Average time, Standard deviation
- **Specific Racer**: Average win time, Raw race data table
- **League Races**: Best win, Worst loss

## Development

### Adding New Components
1. Create component in `src/components/`
2. Use TypeScript interfaces for props
3. Follow PrimeVue styling patterns
4. Register in `main.ts` if used globally

### Adding New API Endpoints
1. Add service function in `src/services/`
2. Create composable in `src/composables/`
3. Add configuration in `src/config.ts`
4. Update types in `src/types.ts`

### Styling
- Uses Tailwind CSS with PrimeUI integration
- Paper Mario color scheme defined in CSS variables
- Custom PrimeVue component styling in `style.css`
- Scoped component styles where needed

## Migration from Original Client

This Vue client replaces the original vanilla TypeScript implementation with:

- **Better State Management**: Vue reactivity vs manual DOM manipulation
- **Component Architecture**: Reusable Vue components vs inline HTML
- **Data Fetching**: TanStack Vue Query vs manual fetch calls
- **Type Safety**: Full TypeScript integration
- **Modern Tooling**: Vite, Vue 3, modern build pipeline

## Troubleshooting

### Common Issues

**API Connection Errors**
- Ensure the statistics API server is running on port 3000
- Check CORS settings if accessing from different domain

**Build Errors**
- Verify Node.js version (requires 20.19+ or 22.12+)
- Clear node_modules and reinstall dependencies

**Styling Issues**
- Ensure PaperMarioFont.ttf is in the public directory
- Check that Tailwind CSS is properly configured

## Contributing

1. Follow TypeScript strict mode
2. Use Vue 3 Composition API
3. Add proper error handling
4. Include loading states
5. Maintain Paper Mario theme consistency
6. Write self-documenting code