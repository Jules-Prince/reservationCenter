import { component$ } from '@builder.io/qwik';

export default component$(() => {
  return (
    <div class="container">
      <h1>Welcome to the Reservation System!</h1>
      <p>
        <a href="/availabilities">View All Availabilities</a>
      </p>
    </div>
  );
});