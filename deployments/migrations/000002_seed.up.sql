BEGIN;

INSERT INTO destinations (id, name) VALUES 
    ('92d8a3e4-76dd-492a-98b7-2058d60c08e4', 'Mars'), 
    ('e308a796-0705-4645-bc00-6e9975203bd9', 'Moon'), 
    ('f3df4162-2a7b-4886-b657-7f78081ab4d1', 'Pluto'),
    ('01c8f5be-a9a5-464a-a038-2da90648dc4d','Asteroid Belt'),
    ('4bab4ddc-599d-4765-a481-44cf059f5733', 'Europa'),
    ('550d5202-9da7-42dc-9a66-a4f2e4b7ea56', 'Titan'),
    ('ef495232-0d9b-4a37-a00c-db856e7b367d', 'Ganymede');

INSERT INTO trips_schedule (destination_id, date) VALUES 
    ('92d8a3e4-76dd-492a-98b7-2058d60c08e4', '2022-06-10'), 
    ('e308a796-0705-4645-bc00-6e9975203bd9', '2022-06-11'), 
    ('f3df4162-2a7b-4886-b657-7f78081ab4d1', '2022-06-12'),
    ('01c8f5be-a9a5-464a-a038-2da90648dc4d','2022-06-13'),
    ('4bab4ddc-599d-4765-a481-44cf059f5733', '2022-06-14'),
    ('550d5202-9da7-42dc-9a66-a4f2e4b7ea56', '2022-06-15'),
    ('ef495232-0d9b-4a37-a00c-db856e7b367d', '2022-06-16'),
    ('92d8a3e4-76dd-492a-98b7-2058d60c08e4', '2022-06-17'), 
    ('e308a796-0705-4645-bc00-6e9975203bd9', '2022-06-18'), 
    ('f3df4162-2a7b-4886-b657-7f78081ab4d1', '2022-06-19'),
    ('01c8f5be-a9a5-464a-a038-2da90648dc4d','2022-06-20'),
    ('4bab4ddc-599d-4765-a481-44cf059f5733', '2022-06-21'),
    ('550d5202-9da7-42dc-9a66-a4f2e4b7ea56', '2022-06-22'),
    ('ef495232-0d9b-4a37-a00c-db856e7b367d', '2022-06-23'),
    ('01c8f5be-a9a5-464a-a038-2da90648dc4d','2022-06-24'),
    ('4bab4ddc-599d-4765-a481-44cf059f5733', '2022-06-25'),
    ('550d5202-9da7-42dc-9a66-a4f2e4b7ea56', '2022-06-26'),
    ('ef495232-0d9b-4a37-a00c-db856e7b367d', '2022-06-27'),
    ('92d8a3e4-76dd-492a-98b7-2058d60c08e4', '2022-06-28'), 
    ('e308a796-0705-4645-bc00-6e9975203bd9', '2022-06-29'), 
    ('f3df4162-2a7b-4886-b657-7f78081ab4d1', '2022-06-30');

INSERT INTO launchpads (id, name, full_name) VALUES 
    ('5e9e4501f5090910d4566f83', 'VAFB SLC 3W', 'Vandenberg Space Force Base Space Launch Complex 3W'),
    ('5e9e4502f509092b78566f87', 'VAFB SLC 4E', 'Vandenberg Space Force Base Space Launch Complex 4E'),
    ('5e9e4501f509094ba4566f84', 'CCSFS SLC 40', 'Cape Canaveral Space Force Station Space Launch Complex 40'),
    ('5e9e4502f5090927f8566f85', 'STLS', 'SpaceX South Texas Launch Site'),
    ('5e9e4502f5090995de566f86', 'Kwajalein Atoll', 'Kwajalein Atoll Omelek Island'),
    ('5e9e4502f509094188566f88', 'KSC LC 39A', 'Kennedy Space Center Historic Launch Complex 39A');

COMMIT;