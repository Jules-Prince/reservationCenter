import { component$, useResource$, Resource } from '@builder.io/qwik';
import { getAvailabilities } from '~/services/api';

export default component$(() => {
  // Fetch availabilities using Qwik's Resource system
  const availabilitiesResource = useResource$(() => getAvailabilities());

  return (
    <div class="container">
      <h1>Availabilities</h1>
      <Resource
        value={availabilitiesResource}
        onPending={() => <p>Loading availabilities...</p>}
        onRejected={(error) => <p>Error: {error.message}</p>}
        onResolved={(availabilities) => (
          <ul>
            {availabilities.map((availability: any, index: number) => (
              <li key={index}>
                <strong>Date:</strong> {availability.date}, 
                <strong> Time:</strong> {availability.startTime} - {availability.endTime}, 
                <strong> Booked:</strong> {availability.isBooked ? 'Yes' : 'No'}
              </li>
            ))}
          </ul>
        )}
      />
    </div>
  );
});
