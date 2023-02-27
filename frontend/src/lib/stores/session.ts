import * as store from "svelte/store";
import * as qg from "$lib/qg";

export const event = store.writable<qg.Event>();

export const session = store.readable<qg.Session>(undefined, (set) => {
  const f = (ev: CustomEvent<qg.Event>) => event.set(ev.detail);
  const s = qg.Session.newLocal();
  s.addEventListener("event", f);

  set(s);

  return () => {
    s.removeEventListener("event", f);
    s.close();
  };
});
